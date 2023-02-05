import { Entity, BaseEntity, Column, PrimaryGeneratedColumn } from 'typeorm';
import { DictionaryType } from './dictionary.model';

@Entity({ name: 'Dictionaries' })
export class Dictionary extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  userId: number;

  @Column()
  type: DictionaryType;

  @Column()
  inputLang: string;

  @Column()
  inputLocaleId: string;

  @Column()
  outputLang: string;

  @Column()
  outputLocaleId: string;

  @Column()
  name: string;

  @Column()
  description: string;

  @Column()
  level: number;

  @Column()
  price: number;

  @Column()
  version: number;

  @Column()
  created: Date;

  @Column()
  updated: Date;

  @Column()
  deleted: Date;
}
