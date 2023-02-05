import { Repository } from 'typeorm';
import { CustomRepository } from '#typeorm-ex/typeorm-ex.decorator';
import { Word } from './word.entity';

@CustomRepository(Word)
export class WordRepository extends Repository<Word> {
  async readWords(): Promise<Word[]> {
    return await this.find();
  }

  async readWordsByDicId(dicId: number) {
    return await this.findBy({ dicId });
  }

  async readWord(id: number) {
    return await this.findOneBy({ id });
  }
}
