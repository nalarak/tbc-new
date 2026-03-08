import * as InputHelpers from '../components/input_helpers.js';
import { HunterOptions_PetType as PetType } from '../proto/hunter.js';
import { ActionId } from '../proto_utils/action_id.js';
import { HunterSpecs } from '../proto_utils/utils.js';

export function makePetTypeInputConfig<SpecType extends HunterSpecs>(): InputHelpers.TypedIconEnumPickerConfig<any, PetType> {
	return InputHelpers.makeClassOptionsEnumIconInput<SpecType, PetType>({
		extraCssClasses: ['pet-type-picker'],
		fieldName: 'petType',
		numColumns: 4,
		values: [
			{ value: PetType.PetNone, actionId: ActionId.fromPetName(''), tooltip: 'No Pet' },
			{ value: PetType.Bat, actionId: ActionId.fromPetName('Bat'), tooltip: 'Bat' },
			{ value: PetType.Bear, actionId: ActionId.fromPetName('Bear'), tooltip: 'Bear' },
			{ value: PetType.Cat, actionId: ActionId.fromPetName('Cat'), tooltip: 'Cat' },
			{ value: PetType.Crab, actionId: ActionId.fromPetName('Crab'), tooltip: 'Crab' },
			{ value: PetType.Owl, actionId: ActionId.fromPetName('Owl'), tooltip: 'Owl' },
			{ value: PetType.Raptor, actionId: ActionId.fromPetName('Raptor'), tooltip: 'Raptor' },
			{ value: PetType.Ravager, actionId: ActionId.fromPetName('Ravager'), tooltip: 'Ravager' },
			{ value: PetType.WindSerpent, actionId: ActionId.fromPetName('Wind Serpent'), tooltip: 'Wind Serpent' },
		],
	});
}
