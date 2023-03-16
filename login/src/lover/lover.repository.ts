import { Repository } from 'typeorm';
import { CustomRepository } from '#typeorm-ex/typeorm-ex.decorator';
import { Lover } from './lover.entity';

@CustomRepository(Lover)
export class LoverRepository extends Repository<Lover> {
  async readLoverByUserId(userId: number): Promise<Lover> {
    return await this.findOneBy({ userId });
  }
}
