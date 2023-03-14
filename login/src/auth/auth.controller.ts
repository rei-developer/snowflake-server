import {
  Controller,
  UseGuards,
  UsePipes,
  ValidationPipe,
  Get,
  Post,
  Delete,
  Body,
} from '@nestjs/common';
import { AuthService } from './auth.service';
import { FirebaseAuthGuard } from '#firebase/firebase-auth.guard';
import { JwtAuthGuard } from './guard/jwt-auth.guard';
import { RequestHeader } from '#common/decorator/request-header.decorator';
import { AuthHeaderRequestDto } from './dto/request/auth-header.request.dto';
import { RegisterRequestDto } from './dto/request/register.request.dto';
import { VerifyResponseDto } from './dto/response/verify.response.dto';
import { UserResponseDto } from '#user/dto/response/user.response.dto';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Get('/verify')
  @UseGuards(FirebaseAuthGuard)
  @UsePipes(ValidationPipe)
  verify(
    @RequestHeader(AuthHeaderRequestDto) authHeaderDto: AuthHeaderRequestDto,
  ): Promise<VerifyResponseDto> {
    return this.authService.verify(authHeaderDto);
  }

  @Get('/verify/custom')
  @UseGuards(JwtAuthGuard)
  @UsePipes(ValidationPipe)
  verifyCustom(
    @RequestHeader(AuthHeaderRequestDto) authHeaderDto: AuthHeaderRequestDto,
  ): Promise<UserResponseDto> {
    return this.authService.verifyCustom(authHeaderDto);
  }

  @Post('/register')
  @UseGuards(JwtAuthGuard)
  @UsePipes(ValidationPipe)
  register(
    @RequestHeader(AuthHeaderRequestDto) authHeaderDto: AuthHeaderRequestDto,
    @Body() registerDto: RegisterRequestDto,
  ): Promise<UserResponseDto> {
    return this.authService.register(authHeaderDto, registerDto);
  }
}
