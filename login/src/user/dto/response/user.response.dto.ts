import { ApiProperty } from '@nestjs/swagger';
import { User } from '#user/user.entity';

export class UserResponseDto {
  constructor(user: User | null) {
    this.user = user;
  }

  @ApiProperty()
  user: User | null;
}
