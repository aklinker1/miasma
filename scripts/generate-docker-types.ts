import YAML from 'yaml';
import { IJsonSchema, OpenAPIV2 } from 'openapi-types';
import CodeBlockWriter from 'code-block-writer';
import { writeFileSync } from 'node:fs';

const SPEC_URL = 'https://docs.docker.com/engine/api/v1.41.yaml';
const writer = new CodeBlockWriter({ indentNumberOfSpaces: 2 });

async function main() {
  const schema: OpenAPIV2.Document = await fetch(SPEC_URL)
    .then(res => res.text())
    .then(YAML.parse);

  writeNamespace('Docker', () => {
    // Generate defintion types
    Object.entries(schema.definitions ?? {}).forEach(([name, model]) => {
      writeTypeDeclaration(name, model);
    });

    // Generate Response Types
    Object.entries(schema.paths).forEach(([path, requests]) => {
      Object.entries(requests).forEach(([method, request]) => {
        if (typeof request !== 'object') return;

        const op = request as OpenAPIV2.OperationObject;
        const opName = op.operationId;
        if (opName == null) throw Error('No operation name: ' + JSON.stringify(op, null, 2));

        Object.entries(op.responses).forEach(([statusCode, response]) => {
          if (response == null) return;

          const typeName = `${titleCase(method)}${opName}Response${statusCode}`;
          if ('$ref' in response) {
            writeTypeDeclaration(typeName, response);
          } else if (response.schema) {
            const docs: string[] = [`\`${method.toUpperCase()} ${path}\``];
            if (op.description) docs.push(op.description.trim());
            if (response.description)
              docs.push(`Code ${statusCode}: ${response.description.trim()}`);
            writeJsdoc(docs.join('\n\n'));
            writeTypeDeclaration(typeName, response.schema);
          }
        });
      });
    });
  });

  const outputFile = 'types/docker.gen.d.ts';
  writeFileSync('types/docker.gen.d.ts', writer.toString());
  console.log(`\n\x1b[32m✔ Generated \x1b[1m${outputFile}\x1b[0m\n`);
}

function writeJsdoc(message?: string) {
  if (!message) return;

  const lines = [
    '/**',
    ...message
      .trim()
      .split('\n')
      .map(line => ` * ${line}`),
    ' */',
  ];
  writer.writeLine(lines.join('\n'));
}

function writeNamespace(name: string, cb: () => void) {
  writer.write(`namespace ${name} `).inlineBlock(cb).newLineIfLastNot().newLine();
}

function writeTypeDeclaration(typeName: string, schema: OpenAPIV2.SchemaObject) {
  writer.write(`type ${typeName} = `);
  writeType(schema);
  writer.write(';').newLineIfLastNot().newLine();
}

function writeType(schema: OpenAPIV2.SchemaObject | IJsonSchema): void {
  // Use Ref name
  if (schema.$ref) {
    const refName = schema.$ref.replace('#/definitions/', '');
    writer.write(`${refName}`);
  } else if (schema.allOf) {
    schema.allOf.forEach((childSchema, i) => {
      writer.conditionalWrite(i !== 0, ' & ');
      writeType(childSchema);
    });
  } else if (schema.oneOf) {
    schema.oneOf.forEach((childSchema, i) => {
      writer.conditionalWrite(i !== 0, ' | ');
      writeType(childSchema);
    });
  } else if (schema.type === 'array' && schema.items) {
    // Wrap arrays in Array<...>
    writer.write('Array<');
    writeType(schema.items);
    writer.write('>');
    return;
  } else if (schema.enum) {
    if (schema.type === 'string') {
      writer.write(schema.enum.map(str => `"${str}"`).join(' | '));
    } else if (schema.type === 'integer') {
      writer.write(schema.enum.map(value => `${value}`).join(' | '));
    } else {
      console.warn('\x1b[33mUnsupported enum type\x1b[0m', schema);
      writer.write('unknown');
    }
  } else if (typeof schema.additionalProperties === 'boolean') {
    writer.write('Record<string, any>');
  } else if (typeof schema.additionalProperties === 'object') {
    writer.write('Record<string, ');
    writeType(schema.additionalProperties);
    writer.write('>');
  } else if (schema.type === 'object') {
    // Write object type
    writer.inlineBlock(() => {
      Object.entries(schema.properties ?? {}).forEach(([propertyName, property]) => {
        writeJsdoc(property.description);
        writer.write(`'${propertyName}'`);
        writer.conditionalWrite(!property.required, '?').write(': ');
        writeType(property);
        writer.write(';').newLineIfLastNot();
      });
    });
  } else if (schema.type === 'boolean' || schema.type === 'string') {
    // Write the type name
    writer.write(schema.type);
  } else if (schema.type === 'integer' || schema.type === 'number') {
    // Write number
    writer.write('number');
  } else if (Array.isArray(schema.type)) {
    // I don't know what this should look like
    throw Error('Array of types not supported: ' + JSON.stringify(schema, null, 2));
  } else {
    console.warn('Unknown type:', schema);
    writer.write(schema.type ?? 'unknown');
  }
}

function titleCase(str: string): string {
  const [first, ...rest] = str;
  return [first.toUpperCase(), ...rest].join('');
}

main();
