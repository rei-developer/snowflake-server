import { ApiProperty } from '@nestjs/swagger';

export class VerifyCustomDto {
  @ApiProperty()
  uid: string | null;
}
