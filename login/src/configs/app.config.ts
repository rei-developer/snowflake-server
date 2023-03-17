export enum LogLevel {
  ERROR = 'error',
  WARN = 'warn',
  INFO = 'info',
  HTTP = 'http',
  VERBOSE = 'verbose',
  DEBUG = 'debug',
  SILLY = 'silly',
}

export const appConfig = {
  name: 'Snowflake Server',
  port: 10000,
  globalPrefix: 'v1',
  rateLimit: {
    ttl: 30,
    limit: 100,
  },
  log: {
    level: LogLevel.DEBUG,
    dir: __dirname + '/../../logs',
    maxFiles: 30,
    zippedArchive: true,
  },
};
