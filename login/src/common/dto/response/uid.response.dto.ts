import { ApiProperty } from '@nestjs/swagger';

export class UidResponseDto {
  @ApiProperty()
  uid: string;
}
