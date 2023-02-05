import { IsNotEmpty, IsNumber } from 'class-validator';

export class GetWordsByDicIdDto {
  @IsNotEmpty()
  @IsNumber()
  dicId: number;
}
