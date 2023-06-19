import DockerModem from 'docker-modem';

export default defineEventHandler(async event => {
  const app = useExtendedNitroApp();
  const user = await app.auth.getUser(event);
  if (user == null) {
    setResponseStatus(event, 403);
    return 'Authentication not sent or invalid';
  }

  // GET requests can't read the body, so return undefined
  const body = await readBody(event).catch(() => undefined);
  const socketPath = import.meta.env.VITE_DOCKER_SOCKET ?? '/var/run/docker.sock';
  const modem = new DockerModem({ socketPath });
  const query = getQuery(event);

  // Including a ? is required to make query params work, and does not break when there are not
  // query params
  const path = `/${event.context.params?.slug}?`;
  try {
    const response = await new Promise((resolve, reject) => {
      const options = {
        path,
        headers: event.req.headers,
        method: event.req.method,
        options: { ...body, ...query },
        statusCodes: { 200: true, 201: true, 204: true },
      };
      modem.dial(options, (err, obj) => {
        if (err != null) reject(err);
        else resolve(obj);
      });
    });
    return response;
  } catch (err) {
    const { statusCode } = err as any;
    event.res.statusCode = statusCode;
    return err;
  }
});
