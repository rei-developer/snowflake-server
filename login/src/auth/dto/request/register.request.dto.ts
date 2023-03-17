import {
  IsNotEmpty,
  IsString,
  Matches,
  IsInt,
  Min,
  Max,
} from 'class-validator';

export class RegisterRequestDto {
  @IsNotEmpty()
  @IsString()
  @Matches(/^([A-Za-z]{2,6}|[가-힣]{2,6})$/)
  readonly name: string;

  @IsInt()
  @Min(1)
  @Max(2)
  readonly sex: number;

  @IsInt()
  @Min(1)
  @Max(20)
  readonly nation: number;
}
