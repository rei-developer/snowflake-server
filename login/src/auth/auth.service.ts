import { Injectable, ForbiddenException, Inject } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { FirebaseAuthStrategy } from '#firebase/firebase-auth.strategy';
import { AuthStrategy } from './auth.strategy';
import { JwtService } from '@nestjs/jwt';
import { UserRepository } from '#user/user.repository';
import { UserService } from '#user/user.service';
import { AuthHeaderDto } from './dto/request/auth-header.dto';
import { VerifyDto } from './dto/response/verify.dto';
import { VerifyCustomDto } from './dto/response/verify-custom.dto';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class AuthService {
  constructor(
    private readonly firebaseAuthStrategy: FirebaseAuthStrategy,
    private readonly authStrategy: AuthStrategy,
    private readonly jwtService: JwtService,
    @InjectRepository(UserRepository)
    private readonly userRepository: UserRepository,
    @Inject(UserService)
    private readonly userService: UserService,
  ) {}

  async verify({ authModel: { idToken } }: AuthHeaderDto): Promise<VerifyDto> {
    try {
      const { uid } = await this.firebaseAuthStrategy.validate(idToken);
      const customToken = this.jwtService.sign({ uid });
      return { uid, idToken, customToken };
    } catch (err) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
  }

  async verifyCustom({
    authModel: { idToken: token },
  }: AuthHeaderDto): Promise<VerifyCustomDto> {
    return await this.authStrategy.validate(token);
  }

  async register({
    authModel: { authType: type, idToken },
  }: AuthHeaderDto): Promise<boolean> {
    const { uid, email } = await this.firebaseAuthStrategy.validate(idToken);
    const user = await this.userRepository.readUserByUId(uid);
    const userModel = {
      uid,
      type,
      email,
    };
    return user
      ? await this.userService.patchUser(uid, userModel)
      : await this.userService.addUser(userModel);
  }

  async withdraw({
    authModel: { idToken: token },
  }: AuthHeaderDto): Promise<boolean> {
    try {
      const { uid } = await this.authStrategy.validate(token);
      if (!uid) {
        return false;
      }
      const { id } = await this.userRepository.readUserByUId(uid);
      return await this.userRepository.deleteUser(id);
    } catch (err) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
  }
}
