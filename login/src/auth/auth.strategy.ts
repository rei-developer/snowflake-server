import { Injectable, Inject } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy, ExtractJwt } from 'passport-firebase-jwt';
import { JwtService } from '@nestjs/jwt';
import { CustomAuthScheme } from './auth.model';
import { appConfig } from '#configs/app.config';
import { secretKeyConfig } from '#configs/secret-key.config';

const {
  jwt: { defaultStrategy },
} = appConfig;

@Injectable()
export class AuthStrategy extends PassportStrategy(Strategy, defaultStrategy) {
  constructor(
    @Inject(JwtService)
    private readonly jwtService: JwtService,
  ) {
    super({
      secretOrKey: secretKeyConfig.secretKey,
      jwtFromRequest: ExtractJwt.fromAuthHeaderWithScheme(CustomAuthScheme),
    });
  }

  async validate(token: string): Promise<{ uid: string | null }> {
    try {
      return await this.jwtService.verifyAsync(token);
    } catch (_) {
      return { uid: null };
    }
  }
}
