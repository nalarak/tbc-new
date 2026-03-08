import * as PresetUtils from '../../core/preset_utils';
import { APLRotation_Type as APLRotationType } from '../../core/proto/apl.js';
import { Class, ConsumesSpec, Debuffs, IndividualBuffs, PartyBuffs, Profession, PseudoStat, RaidBuffs, Stat, TristateEffect } from '../../core/proto/common';
import { HunterOptions_PetType as PetType, Hunter_Options as HunterOptions, HunterOptions_Ammo, HunterOptions_QuiverBonus } from '../../core/proto/hunter';
import { SavedTalents } from '../../core/proto/ui';
import { Stats } from '../../core/proto_utils/stats';
import { defaultRaidBuffMajorDamageCooldowns } from '../../core/proto_utils/utils';
import DefaultAPL from './apls/default.apl.json';
import P1_2HGear from './gear_sets/p1_2h.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const DEFAULT_APL = PresetUtils.makePresetAPLRotation('Default', DefaultAPL);

export const P1_2H_GEARSET = PresetUtils.makePresetGear('P1 - 2H', P1_2HGear);

export const P1_EP_PRESET = PresetUtils.makePresetEpWeights(
	'P1',
	Stats.fromMap(
		{
			[Stat.StatAgility]: 1,
			[Stat.StatStrength]: 0.06,
			[Stat.StatIntellect]: 0.01,
			[Stat.StatAttackPower]: 0.06,
			[Stat.StatRangedAttackPower]: 0.4,
			[Stat.StatMeleeHitRating]: 0.12,
			[Stat.StatMeleeCritRating]: 0.92,
			[Stat.StatMeleeHasteRating]: 0.788,
			[Stat.StatArmorPenetration]: 0.16,
		},
		{
			[PseudoStat.PseudoStatRangedDps]: 1.75,
		},
	),
);

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/wotlk/talent-calc and copy the numbers in the url.

export const BMTalents = {
	name: 'BM',
	data: SavedTalents.create({
		talentsString: '512002005250122431051-0505201205',
	}),
};
export const SVTalents = {
	name: 'SV',
	data: SavedTalents.create({
		talentsString: '502-0550201205-333200022003223005103',
	}),
};

export const DefaultOptions = HunterOptions.create({
	classOptions: {
		ammo: HunterOptions_Ammo.WardensArrow,
		quiverBonus: HunterOptions_QuiverBonus.Speed15,
		petType: PetType.Ravager,
		petUptime: 1,
		petSingleAbility: true,
	},
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	unleashedRage: true,
});

export const DefaultPartyBuffs = PartyBuffs.create({
	ferociousInspiration: 2,
	braidedEterniumChain: true,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	windfuryTotem: TristateEffect.TristateEffectImproved,
	battleShout: TristateEffect.TristateEffectImproved,
	leaderOfThePack: TristateEffect.TristateEffectImproved,
	trueshotAura: true,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	...defaultRaidBuffMajorDamageCooldowns(Class.ClassWarrior),
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
});

export const DefaultDebuffs = Debuffs.create({
	improvedSealOfTheCrusader: true,
	misery: true,
	bloodFrenzy: true,
	giftOfArthas: true,
	mangle: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: TristateEffect.TristateEffectImproved,
	sunderArmor: true,
	curseOfRecklessness: true,
	huntersMark: TristateEffect.TristateEffectImproved,
	exposeWeaknessUptime: 0.9,
	exposeWeaknessHunterAgility: 1080,
});

export const DefaultConsumables = ConsumesSpec.create({
	flaskId: 22854, // Flask of Relentless Assault
	foodId: 27659, // Warp Burger
	potId: 22838, // Haste Potion
	conjuredId: 5512,
	explosiveId: 30217,
	drumsId: 351355,
	petScrollAgi: true,
	petScrollStr: true,
	superSapper: true,
	goblinSapper: true,
	scrollAgi: true,
	scrollStr: true,
});

export const OtherDefaults = {
	distanceFromTarget: 7,
	iterationCount: 25000,
	profession1: Profession.Engineering,
	profession2: Profession.Blacksmithing,
};
