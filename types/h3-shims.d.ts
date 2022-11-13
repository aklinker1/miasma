type H3Error<T> = {
  message: string;
  statusCode: number;
  data: {
    url: string;
    statusCode: 503;
    statusMessage: string;
    stack?: string;
  } & T;
};
