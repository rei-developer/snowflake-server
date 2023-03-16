import { IsString } from 'class-validator';

export class LangRequestDto {
  @IsString()
  lang?: string = 'en';
}
