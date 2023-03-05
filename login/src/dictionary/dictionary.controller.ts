import { Controller, Get, Param } from '@nestjs/common';
import { DictionaryService } from './dictionary.service';
import { GetDictionaryDto } from './dto/request/get-dictionary.dto';
import { Dictionary } from './dictionary.entity';

@Controller('dictionary')
export class DictionaryController {
  constructor(private readonly dictionaryService: DictionaryService) {}

  @Get('/list')
  getDictionaries(): Promise<Dictionary[]> {
    return this.dictionaryService.getDictionaries();
  }

  @Get('/:id')
  getDictionary(
    @Param() getDictionaryDto: GetDictionaryDto,
  ): Promise<Dictionary> {
    return this.dictionaryService.getDictionary(getDictionaryDto);
  }
}
