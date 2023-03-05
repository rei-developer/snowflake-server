import { Injectable } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';
import { appConfig } from '#configs/app.config';

const {
  jwt: { defaultStrategy },
} = appConfig;

@Injectable()
export class JwtAuthGuard extends AuthGuard(defaultStrategy) {}
