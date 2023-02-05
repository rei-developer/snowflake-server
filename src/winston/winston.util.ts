import { utilities, WinstonModule } from 'nest-winston';
import { dailyOptions } from './winston.const';
import * as winstonDaily from 'winston-daily-rotate-file';
import * as winston from 'winston';
import { appConfig, LogLevel } from '#configs/app.config';

const {
  name,
  log: { level },
} = appConfig;
const { NODE_ENV: env } = process.env;

export const winstonLogger = WinstonModule.createLogger({
  transports: [
    new winston.transports.Console({
      level,
      format:
        env === 'production'
          ? winston.format.simple()
          : winston.format.combine(
              winston.format.timestamp(),
              utilities.format.nestLike(name, {
                colors: true,
                prettyPrint: true,
              }),
            ),
    }),
    new winstonDaily(dailyOptions(LogLevel.INFO)),
    new winstonDaily(dailyOptions(LogLevel.WARN)),
    new winstonDaily(dailyOptions(LogLevel.ERROR)),
  ],
});
