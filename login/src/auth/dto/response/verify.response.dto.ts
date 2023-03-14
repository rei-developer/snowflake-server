import { ApiProperty } from '@nestjs/swagger';

export class VerifyResponseDto {
  @ApiProperty()
  uid: string;

  @ApiProperty()
  customToken: string;
}
