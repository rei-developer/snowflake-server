import * as fs from 'fs';
import * as yaml from 'yaml';
import { Injectable } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';

const { defaultStrategy } = yaml.parse(
  fs.readFileSync('config.yaml', 'utf8'),
).jwt;

@Injectable()
export class JwtAuthGuard extends AuthGuard(defaultStrategy) {}
