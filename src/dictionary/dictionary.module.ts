import { Module } from '@nestjs/common';
import { TypeOrmExModule } from '#typeorm-ex/typeorm-ex.module';
import { DictionaryRepository } from './dictionary.repository';
import { DictionaryController } from './dictionary.controller';
import { DictionaryService } from './dictionary.service';

@Module({
  imports: [TypeOrmExModule.forCustomRepository([DictionaryRepository])],
  controllers: [DictionaryController],
  providers: [DictionaryService],
})
export class DictionaryModule {}
