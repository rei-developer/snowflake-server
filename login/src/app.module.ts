import { Module, NestModule, MiddlewareConsumer } from '@nestjs/common';
import { ThrottlerModule, ThrottlerGuard } from '@nestjs/throttler';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AuthModule } from '#auth/auth.module';
import { UserModule } from '#user/user.module';
import { AppController } from '#app.controller';
import { AppService } from '#app.service';
import { LoggerMiddleware } from '#common/middleware/logger.middleware';
import { APP_GUARD } from '@nestjs/core';
import { appConfig } from '#configs/app.config';
import { typeORMConfig } from '#configs/typeorm.config';

const {
  rateLimit: { ttl, limit },
} = appConfig;

@Module({
  imports: [
    ThrottlerModule.forRoot({
      ttl,
      limit,
    }),
    TypeOrmModule.forRoot(typeORMConfig),
    AuthModule,
    UserModule,
  ],
  controllers: [AppController],
  providers: [
    AppService,
    {
      provide: APP_GUARD,
      useClass: ThrottlerGuard,
    },
  ],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(LoggerMiddleware).forRoutes('*');
  }
}
