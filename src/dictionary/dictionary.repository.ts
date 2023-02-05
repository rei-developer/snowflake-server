import { Repository } from 'typeorm';
import { CustomRepository } from '#typeorm-ex/typeorm-ex.decorator';
import { Dictionary } from './dictionary.entity';

@CustomRepository(Dictionary)
export class DictionaryRepository extends Repository<Dictionary> {
  async readDictionaries(): Promise<Dictionary[]> {
    return await this.find();
  }

  async readDictionary(id: number): Promise<Dictionary> {
    return await this.findOneBy({ id });
  }
}
