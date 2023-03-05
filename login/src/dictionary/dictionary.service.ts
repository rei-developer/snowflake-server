import { Injectable, ForbiddenException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DictionaryRepository } from './dictionary.repository';
import { GetDictionaryDto } from '#dictionary/dto/request/get-dictionary.dto';
import { Dictionary } from './dictionary.entity';
import { ExceptionErrorMessage } from '#common/const/exception-error-message.const';

@Injectable()
export class DictionaryService {
  constructor(
    @InjectRepository(DictionaryRepository)
    private readonly dictionaryRepository: DictionaryRepository,
  ) {}

  async getDictionaries(): Promise<Dictionary[]> {
    const dictionaries = await this.dictionaryRepository.readDictionaries();
    if (!dictionaries) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    return dictionaries;
  }

  async getDictionary({ id }: GetDictionaryDto): Promise<Dictionary> {
    const dictionary = await this.dictionaryRepository.readDictionary(id);
    if (!dictionary) {
      throw new ForbiddenException(ExceptionErrorMessage.DOES_NOT_EXIST);
    }
    return dictionary;
  }
}
