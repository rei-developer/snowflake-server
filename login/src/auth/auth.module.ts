import * as fs from 'fs';
import * as yaml from 'yaml';
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

const { defaultStrategy, secretKey, expiresIn } = yaml.parse(
  fs.readFileSync('config.yaml', 'utf8'),
).jwt;

@Module({
  imports: [
    FirebaseAuthModule,
    PassportModule.register({ defaultStrategy }),
    JwtModule.register({
      secret: secretKey,
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
