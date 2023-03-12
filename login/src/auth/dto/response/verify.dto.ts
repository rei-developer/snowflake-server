import { ApiProperty } from '@nestjs/swagger';

export class VerifyDto {
  @ApiProperty()
  uid: string;

  @ApiProperty()
  customToken: string;
}
