import { H3Event } from 'h3';

export default function (template: string | undefined): Auth {
  let auth: Auth;

  if (template?.startsWith('token:')) auth = defineTokenAuth(template);
  else if (template?.startsWith('basic:')) auth = defineBasicAuth(template);
  else {
    if (template != null) {
      console.warn(
        "auth env variable does not start with a known prefix ('token:...' or 'basic:...'). See https://miasma.aklinker1.io/authentication for more details.",
      );
    }
    auth = defineNoAuth();
  }

  return {
    type: auth.type,
    async getUser(event) {
      const user = await auth.getUser(event);
      //. If we're not logged in, sleep randomly between 3 seconds to slow down reqeusts
      if (user == null) await new Promise(res => setTimeout(res, Math.random() * 3000));
      return user;
    },
  };
}

export interface User {
  username: string;
}

export interface Auth {
  type: string;
  /**
   * Returns the valid user if authentication is correct.
   */
  getUser(event: H3Event): User | undefined | Promise<User | undefined>;
}

function defineTokenAuth(template: string): Auth {
  const token = template.replace('token:', '');

  return {
    type: 'token',
    getUser(event) {
      const authHeader = getHeader(event, 'Authorization');

      if (authHeader === `Bearer ${token}`) return { username: 'Admin' };
      else return undefined;
    },
  };
}

function defineBasicAuth(template: string): Auth {
  const users = template
    .replace('basic:', '')
    .split('\n')
    .reduce<Record<string, string>>((map, userTemplate) => {
      const [username, password] = userTemplate
        .trim()
        .split(':', 2)
        .map(str => str.trim());
      if (!username || !password) {
        console.warn('Username or password missing, this one user ignored.');
      } else {
        map[username] = password;
      }
      return map;
    }, {});

  return {
    type: 'basic',
    getUser(event) {
      const encodedAuthHeader = getHeader(event, 'Authorization');
      if (!encodedAuthHeader?.startsWith('Basic ')) return undefined;

      const decodedAuthHeader = Buffer.from(
        encodedAuthHeader.replace('Basic ', ''),
        'base64',
      ).toString('utf-8');
      const [username, password] = decodedAuthHeader.split(':', 2);

      if (!username || !password) return undefined;
      if (users[username] !== password) return undefined;

      return { username };
    },
  };
}

function defineNoAuth(): Auth {
  return {
    type: 'none',
    getUser() {
      return { username: 'Admin' };
    },
  };
}
