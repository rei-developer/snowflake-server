import { IsString } from 'class-validator';

export class LangDto {
  @IsString()
  lang?: string = 'en';
}
