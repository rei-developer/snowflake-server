import * as fs from 'fs';
import * as yaml from 'yaml';
import { Injectable, Inject } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { Strategy, ExtractJwt } from 'passport-firebase-jwt';
import { JwtService } from '@nestjs/jwt';
import { CustomAuthScheme } from './auth.model';

const { defaultStrategy, secretKey } = yaml.parse(
  fs.readFileSync('config.yaml', 'utf8'),
).jwt;

@Injectable()
export class AuthStrategy extends PassportStrategy(Strategy, defaultStrategy) {
  constructor(
    @Inject(JwtService)
    private readonly jwtService: JwtService,
  ) {
    super({
      secretOrKey: secretKey,
      jwtFromRequest: ExtractJwt.fromAuthHeaderWithScheme(CustomAuthScheme),
    });
  }

  async validate(token: string): Promise<{ jti: string | null }> {
    try {
      return await this.jwtService.verifyAsync(token);
    } catch (_) {
      return { jti: null };
    }
  }
}
