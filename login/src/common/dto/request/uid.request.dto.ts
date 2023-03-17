import { IsNotEmpty, IsString } from 'class-validator';

export class UidRequestDto {
  @IsNotEmpty()
  @IsString()
  readonly uid: string;
}
