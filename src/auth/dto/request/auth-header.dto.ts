import { Expose, Transform } from 'class-transformer';
import { AuthType, AuthModel } from '#auth/auth.model';

export class AuthHeaderDto {
  @Expose({ name: 'authorization' })
  @Transform(({ value }: { value: string }) => {
    const [authType, idToken] = value.split(' ');
    return {
      authType: AuthType[authType.toUpperCase()],
      idToken,
    };
  })
  authModel: AuthModel;
}
