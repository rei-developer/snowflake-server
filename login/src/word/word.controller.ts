import { Controller, Get, Param } from '@nestjs/common';
import { WordService } from './word.service';
import { GetWordsByDicIdDto } from './dto/request/get-words-by-dic-id.dto';
import { GetWordDto } from './dto/request/get-word.dto';
import { Word } from './word.entity';

@Controller('word')
export class WordController {
  constructor(private readonly wordService: WordService) {}

  @Get('/list')
  getWords(): Promise<Word[]> {
    return this.wordService.getWords();
  }

  @Get('/list/:dicId')
  getWordsByDicId(
    @Param() getWordsByDicIdDto: GetWordsByDicIdDto,
  ): Promise<Word[]> {
    return this.wordService.getWordsByDicId(getWordsByDicIdDto);
  }

  @Get('/:id')
  getWord(@Param() getWordDto: GetWordDto): Promise<Word> {
    return this.wordService.getWord(getWordDto);
  }
}
