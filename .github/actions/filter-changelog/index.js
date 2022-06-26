const core = require("@actions/core");

try {
  // Read inputs
  const input = core.getInput("changelog");
  const inputLines = input.split("\n");
  const scopes = core
    .getInput("scopes")
    .split(",")
    .map((s) => s.trim());

  // Filter Changelog
  const outputLines = [
    "Download via [Docker Hub](https://hub.docker.com/r/aklinker1/miasma)",
    "",
  ];
  let scopedCount = 0;
  inputLines.forEach((line) => {
    // Line isn't a bullet point
    if (!line.startsWith("* ")) return outputLines.push(line);

    // Bullet point starts with the scope prefix
    const matchesScope = !!scopes.find((scope) =>
      line.toLowerCase().startsWith(`* **${scope.toLowerCase()}:**`)
    );
    if (matchesScope) {
      scopedCount++;
      return outputLines.push(line);
    }
  });
  const output = outputLines.join("\n");
  const empty = scopedCount === 0;

  // Set outputs
  console.log({ input, output, empty });
  core.setOutput("changelog", output);
  core.setOutput("empty", empty);
} catch (error) {
  core.setFailed(error.message);
}
