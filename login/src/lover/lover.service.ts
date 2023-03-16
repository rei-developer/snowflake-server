import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { LoverRepository } from './lover.repository';
import { Lover } from './lover.entity';

@Injectable()
export class LoverService {
  constructor(
    @InjectRepository(LoverRepository)
    private readonly loverRepository: LoverRepository,
  ) {}

  async fetchLoverByUserId(userId: number): Promise<Lover> {
    return await this.loverRepository.readLoverByUserId(userId);
  }
}
