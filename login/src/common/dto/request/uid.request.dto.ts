import { IsNotEmpty, IsString } from 'class-validator';

export class UidRequestDto {
  @IsNotEmpty()
  @IsString()
  uid: string;
}
