import { ForbiddenException, Inject, Injectable } from '@nestjs/common';
import { FirebaseAuthStrategy } from '#firebase/firebase-auth.strategy';
import { AuthStrategy } from './auth.strategy';
import { JwtService } from '@nestjs/jwt';
import { LoverService } from '#lover/lover.service';
import { UserService } from '#user/user.service';
import { AuthHeaderRequestDto } from './dto/request/auth-header.request.dto';
import { VerifyResponseDto } from './dto/response/verify.response.dto';
import { RegisterRequestDto } from './dto/request/register.request.dto';
import { SignInResultResponseDto } from './dto/response/common/sign-in-result.response.dto';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class AuthService {
  constructor(
    private readonly firebaseAuthStrategy: FirebaseAuthStrategy,
    private readonly authStrategy: AuthStrategy,
    private readonly jwtService: JwtService,
    @Inject(LoverService)
    private readonly loverService: LoverService,
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
  }: AuthHeaderRequestDto): Promise<SignInResultResponseDto> {
    const { uid } = await this.authStrategy.validate(token);
    const user = await this.userService.fetchUser(uid);
    return new SignInResultResponseDto(uid, !!user, false);
  }

  async register(
    { authModel: { idToken: token } }: AuthHeaderRequestDto,
    { name }: RegisterRequestDto,
  ): Promise<SignInResultResponseDto> {
    const { uid } = await this.authStrategy.validate(token);
    const user = await this.userService.fetchUser(uid);
    const userId = user
      ? user.id
      : await this.userService.addUser({ uid, name });
    if (userId < 1) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    const lover = await this.loverService.fetchLoverByUserId(userId);
    return new SignInResultResponseDto(uid, userId > 0, !!lover);
  }
}
