import * as InputHelpers from '../../core/components/input_helpers';
import { Spec } from '../../core/proto/common';
import { HunterOptions_Ammo, HunterOptions_QuiverBonus } from '../../core/proto/hunter';
import { ActionId } from '../../core/proto_utils/action_id';
import { makePetTypeInputConfig } from '../../core/talents/hunter_pet';
import i18n from '../../i18n/config.js';

export const AmmoInput = () =>
	InputHelpers.makeClassOptionsEnumIconInput<Spec.SpecHunter, HunterOptions_Ammo>({
		fieldName: 'ammo',
		numColumns: 7,
		values: [
			{ value: HunterOptions_Ammo.AmmoNone, tooltip: 'No Ammo' },
			{ actionId: ActionId.fromItemId(31737), value: HunterOptions_Ammo.TimelessArrow },
			{ actionId: ActionId.fromItemId(34581), value: HunterOptions_Ammo.MysteriousArrow },
			{ actionId: ActionId.fromItemId(33803), value: HunterOptions_Ammo.AdamantiteStinger },
			{ actionId: ActionId.fromItemId(30611), value: HunterOptions_Ammo.HalaaniRazorshaft },
			{ actionId: ActionId.fromItemId(28056), value: HunterOptions_Ammo.BlackflightArrow },
			{ actionId: ActionId.fromItemId(31949), value: HunterOptions_Ammo.WardensArrow },
		],
	});

export const QuiverInput = () =>
	InputHelpers.makeClassOptionsEnumIconInput<Spec.SpecHunter, HunterOptions_QuiverBonus>({
		extraCssClasses: ['quiver-picker'],
		fieldName: 'quiverBonus',
		numColumns: 7,
		values: [
			{ color: '82e89d', value: HunterOptions_QuiverBonus.QuiverNone },
			{ actionId: ActionId.fromItemId(18714), value: HunterOptions_QuiverBonus.Speed15 },
			{ actionId: ActionId.fromItemId(2662), value: HunterOptions_QuiverBonus.Speed14 },
			{ actionId: ActionId.fromItemId(8217), value: HunterOptions_QuiverBonus.Speed13 },
			{ actionId: ActionId.fromItemId(7371), value: HunterOptions_QuiverBonus.Speed12 },
			{ actionId: ActionId.fromItemId(3605), value: HunterOptions_QuiverBonus.Speed11 },
			{ actionId: ActionId.fromItemId(3573), value: HunterOptions_QuiverBonus.Speed10 },
		],
	});

export const PetTypeInput = () => makePetTypeInputConfig<Spec.SpecHunter>();

export const PetSingleAbility = () =>
	InputHelpers.makeClassOptionsBooleanInput<Spec.SpecHunter>({
		fieldName: 'petSingleAbility',
		label: i18n.t('settings_tab.other.pet_single_ability.label'),
		labelTooltip: i18n.t('settings_tab.other.pet_single_ability.tooltip'),
	});

export const PetUptime = () =>
	InputHelpers.makeClassOptionsNumberInput<Spec.SpecHunter>({
		fieldName: 'petUptime',
		label: i18n.t('settings_tab.other.pet_uptime.label'),
		labelTooltip: i18n.t('settings_tab.other.pet_uptime.tooltip'),
		percent: true,
	});
