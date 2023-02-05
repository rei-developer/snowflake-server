import { NestFactory } from '@nestjs/core';
import { AppModule } from '#app.module';
import { winstonLogger } from '#winston/winston.util';
import { ValidationPipe, Logger } from '@nestjs/common';
import { ResponseHeaderInterceptor } from '#common/interceptor/response-header.interceptor';
import { appConfig } from '#configs/app.config';
import 'module-alias/register';

const { name, port, globalPrefix } = appConfig;
const { NODE_ENV: env } = process.env;

(async () =>
  (await NestFactory.create(AppModule, { logger: winstonLogger }))
    .useGlobalPipes(
      new ValidationPipe({
        transform: true,
        transformOptions: {
          enableImplicitConversion: true,
        },
      }),
    )
    .useGlobalInterceptors(new ResponseHeaderInterceptor())
    .setGlobalPrefix(globalPrefix)
    .listen(port)
    .then(() => Logger.log(`${name} - ${env} running on port ${port}`)))();
