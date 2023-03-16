import { Module } from '@nestjs/common';
import { TypeOrmExModule } from '#typeorm-ex/typeorm-ex.module';
import { LoverRepository } from './lover.repository';
import { LoverService } from './lover.service';

@Module({
  imports: [TypeOrmExModule.forCustomRepository([LoverRepository])],
  exports: [LoverService],
  providers: [LoverService],
})
export class LoverModule {}
