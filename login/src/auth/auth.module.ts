import { Module } from '@nestjs/common';
import { FirebaseAuthModule } from '#firebase/firebase-auth.module';
import { PassportModule } from '@nestjs/passport';
import { JwtModule } from '@nestjs/jwt';
import { TypeOrmExModule } from '#typeorm-ex/typeorm-ex.module';
import { UserRepository } from '#user/user.repository';
import { LoverModule } from '#lover/lover.module';
import { UserModule } from '#user/user.module';
import { AuthStrategy } from './auth.strategy';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { appConfig } from '#configs/app.config';
import { secretKeyConfig } from '#configs/secret-key.config';

const {
  jwt: { defaultStrategy, expiresIn },
} = appConfig;

@Module({
  imports: [
    FirebaseAuthModule,
    PassportModule.register({ defaultStrategy: defaultStrategy }),
    JwtModule.register({
      secret: secretKeyConfig.secretKey,
      signOptions: { expiresIn },
    }),
    TypeOrmExModule.forCustomRepository([UserRepository]),
    LoverModule,
    UserModule,
  ],
  exports: [PassportModule, AuthStrategy],
  controllers: [AuthController],
  providers: [AuthService, AuthStrategy],
})
export class AuthModule {}
