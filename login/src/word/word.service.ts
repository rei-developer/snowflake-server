import { Injectable, ForbiddenException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { WordRepository } from './word.repository';
import { DictionaryRepository } from '#dictionary/dictionary.repository';
import { GetWordsByDicIdDto } from './dto/request/get-words-by-dic-id.dto';
import { GetWordDto } from './dto/request/get-word.dto';
import { Word } from './word.entity';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class WordService {
  constructor(
    @InjectRepository(WordRepository)
    private readonly wordRepository: WordRepository,
    @InjectRepository(DictionaryRepository)
    private readonly dictionaryRepository: DictionaryRepository,
  ) {}

  async getWords(): Promise<Word[]> {
    const words = await this.wordRepository.readWords();
    if (!words) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    return words;
  }

  async getWordsByDicId({ dicId }: GetWordsByDicIdDto): Promise<Word[]> {
    if (!(await this.dictionaryRepository.readDictionary(dicId))) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    const words = await this.wordRepository.readWordsByDicId(dicId);
    if (!words) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    return words;
  }

  async getWord({ id }: GetWordDto): Promise<Word> {
    const word = await this.wordRepository.readWord(id);
    if (!word) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    return word;
  }
}
