import { IsNotEmpty, IsString } from 'class-validator';

export class RegisterRequestDto {
  @IsNotEmpty()
  @IsString()
  name: string;
}
