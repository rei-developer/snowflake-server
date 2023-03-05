import {
  Controller,
  UseGuards,
  UsePipes,
  ValidationPipe,
  Get,
  Post,
  Delete,
  Query,
  Body,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { FirebaseAuthGuard } from '#firebase/firebase-auth.guard';
import { JwtAuthGuard } from './guard/jwt-auth.guard';
import { RequestHeader } from '#common/decorator/request-header.decorator';
import { AuthHeaderDto } from './dto/request/auth-header.dto';
import { VerifyDto } from './dto/response/verify.dto';
import { VerifyCustomDto } from './dto/response/verify-custom.dto';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Get('/verify')
  @UseGuards(FirebaseAuthGuard)
  @UsePipes(ValidationPipe)
  verify(
    @RequestHeader(AuthHeaderDto) authHeaderDto: AuthHeaderDto,
  ): Promise<VerifyDto> {
    return this.authService.verify(authHeaderDto);
  }

  @Get('/verify/custom')
  @UseGuards(JwtAuthGuard)
  @UsePipes(ValidationPipe)
  verifyCustom(
    @RequestHeader(AuthHeaderDto) authHeaderDto: AuthHeaderDto,
  ): Promise<VerifyCustomDto> {
    return this.authService.verifyCustom(authHeaderDto);
  }

  @Post('/register')
  @UseGuards(FirebaseAuthGuard)
  @UsePipes(ValidationPipe)
  register(
    @RequestHeader(AuthHeaderDto) authHeaderDto: AuthHeaderDto,
  ): Promise<boolean> {
    return this.authService.register(authHeaderDto);
  }

  @Delete('/withdraw')
  @UseGuards(FirebaseAuthGuard)
  @UsePipes(ValidationPipe)
  withdraw(
    @RequestHeader(AuthHeaderDto) authHeaderDto: AuthHeaderDto,
  ): Promise<boolean> {
    return this.authService.withdraw(authHeaderDto);
  }
}
