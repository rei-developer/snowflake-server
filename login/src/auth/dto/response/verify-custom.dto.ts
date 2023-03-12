import { ApiProperty } from '@nestjs/swagger';

export class VerifyCustomDto {
  constructor(uid: string, hasUser: boolean) {
    this.uid = uid;
    this.hasUser = hasUser;
  }

  @ApiProperty()
  uid: string | null;

  @ApiProperty()
  hasUser: boolean;
}
