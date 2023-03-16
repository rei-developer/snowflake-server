import { IsNumber } from 'class-validator';
import { Transform } from 'class-transformer';

export class LimitRequestDto {
  @IsNumber()
  @Transform(({ value }) => parseInt(value))
  limit?: number = 1;
}
