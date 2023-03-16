import { ApiProperty } from '@nestjs/swagger';

export class SignInResultResponseDto {
  constructor(hasUser: boolean, hasLover: boolean) {
    this.hasUser = hasUser;
    this.hasLover = hasLover;
  }

  @ApiProperty()
  hasUser: boolean;

  @ApiProperty()
  hasLover: boolean;
}
