import { Module } from '@nestjs/common';
import { TypeOrmExModule } from '#typeorm-ex/typeorm-ex.module';
import { WordRepository } from './word.repository';
import { DictionaryRepository } from '#dictionary/dictionary.repository';
import { WordController } from './word.controller';
import { WordService } from './word.service';

@Module({
  imports: [
    TypeOrmExModule.forCustomRepository([WordRepository, DictionaryRepository]),
  ],
  controllers: [WordController],
  providers: [WordService],
})
export class WordModule {}
