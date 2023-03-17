import { ApiProperty } from '@nestjs/swagger';
import { UidResponseDto } from '#common/dto/response/uid.response.dto';

export class SignInResultResponseDto extends UidResponseDto {
  constructor(uid: string, hasUser: boolean, hasLover: boolean) {
    super();
    this.uid = uid;
    this.hasUser = hasUser;
    this.hasLover = hasLover;
  }

  @ApiProperty()
  readonly hasUser: boolean;

  @ApiProperty()
  readonly hasLover: boolean;
}
