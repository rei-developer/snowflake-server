import { ApiProperty } from '@nestjs/swagger';
import { UidResponseDto } from '#common/dto/response/uid.response.dto';

export class VerifyResponseDto extends UidResponseDto {
  @ApiProperty()
  readonly customToken: string;
}
