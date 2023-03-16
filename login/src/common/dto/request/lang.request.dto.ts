import { IsString } from 'class-validator';

export class LangRequestDto {
  @IsString()
  readonly lang?: string = 'en';
}
