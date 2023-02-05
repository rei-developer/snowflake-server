import { IsNotEmpty, IsString } from 'class-validator';

export class UIdDto {
  @IsNotEmpty()
  @IsString()
  uid: string;
}
