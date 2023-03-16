import { IsNotEmpty, IsNumber } from 'class-validator';

export class IdRequestDto {
  @IsNotEmpty()
  @IsNumber()
  id: number;
}
