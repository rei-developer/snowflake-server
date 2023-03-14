import { ForbiddenException, Inject, Injectable } from '@nestjs/common';
import { FirebaseAuthStrategy } from '#firebase/firebase-auth.strategy';
import { AuthStrategy } from './auth.strategy';
import { JwtService } from '@nestjs/jwt';
import { UserService } from '#user/user.service';
import { AuthHeaderRequestDto } from './dto/request/auth-header.request.dto';
import { VerifyResponseDto } from './dto/response/verify.response.dto';
import { RegisterRequestDto } from './dto/request/register.request.dto';
import { UserResponseDto } from '#user/dto/response/user.response.dto';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class AuthService {
  constructor(
    private readonly firebaseAuthStrategy: FirebaseAuthStrategy,
    private readonly authStrategy: AuthStrategy,
    private readonly jwtService: JwtService,
    @Inject(UserService)
    private readonly userService: UserService,
  ) {}

  async verify({
    authModel: { idToken },
  }: AuthHeaderRequestDto): Promise<VerifyResponseDto> {
    try {
      const { uid } = await this.firebaseAuthStrategy.validate(idToken);
      const customToken = this.jwtService.sign({ uid });
      return { uid, customToken };
    } catch (err) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
  }

  async verifyCustom({
    authModel: { idToken: token },
  }: AuthHeaderRequestDto): Promise<UserResponseDto> {
    const { uid } = await this.authStrategy.validate(token);
    return await this.userService.fetchUser(uid);
  }

  async register(
    { authModel: { idToken: token } }: AuthHeaderRequestDto,
    { name }: RegisterRequestDto,
  ): Promise<UserResponseDto> {
    const { uid } = await this.authStrategy.validate(token);
    const user = await this.userService.fetchUser(uid);
    if (user.user) return user;
    await this.userService.addUser({ uid, name });
    console.log('유저 생성 완료입니다');
    return await this.userService.fetchUser(uid);
  }
}
