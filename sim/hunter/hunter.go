package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const ThoridalTheStarsFuryItemID = 34334

var TalentTreeSizes = [3]int{21, 20, 24}

const (
	ClassSpellMask_HunterNone uint64 = 0

	// Shots
	ClassSpellMask_HunterAimedShot uint64 = 1 << iota
	ClassSpellMask_HunterArcaneShot
	ClassSpellMask_HunterChimeraShot
	ClassSpellMask_HunterExplosiveShot
	ClassSpellMask_HunterKillShot
	ClassSpellMask_HunterMultiShot
	ClassSpellMask_HunterSteadyShot

	// Strikes
	ClassSpellMask_HunterFlankingStrike
	ClassSpellMask_HunterRaptorStrike
	ClassSpellMask_HunterRaptorStrikeHit
	ClassSpellMask_HunterWyvernStrike

	// Stings
	ClassSpellMask_HunterSerpentSting
	ClassSpellMask_HunterSoFSerpentSting // Serpent Sting from [Strings of Fate]

	// Traps
	ClassSpellMask_HunterExplosiveTrap
	ClassSpellMask_HunterFreezingTrap
	ClassSpellMask_HunterImmolationTrap

	// Other
	ClassSpellMask_HunterCarve
	ClassSpellMask_HunterCarveHit
	ClassSpellMask_HunterMongooseBite
	ClassSpellMask_HunterWingClip
	ClassSpellMask_HunterVolley
	ClassSpellMask_HunterChimeraSerpent
	ClassSpellMask_HunterHuntersMark
	ClassSpellMask_HunterFocusFire

	// Pet Spells
	ClassSpellMask_HunterPetBite
	ClassSpellMask_HunterPetClaw
	ClassSpellMask_HunterPetFlankingStrike
	ClassSpellMask_HunterPetLightningBreath
	ClassSpellMask_HunterPetLavaBreath
	ClassSpellMask_HunterPetScreech
	ClassSpellMask_HunterPetScorpidPoison
	ClassSpellMask_HunterPetBasicAttacks   = ClassSpellMask_HunterPetBite | ClassSpellMask_HunterPetClaw | ClassSpellMask_HunterPetLightningBreath | ClassSpellMask_HunterPetLavaBreath | ClassSpellMask_HunterPetScorpidPoison
	ClassSpellMask_HunterPetSpecialAttacks = ClassSpellMask_HunterPetScreech // TODO: Other specials?

	ClassSpellMask_HunterAll = 1<<iota - 1

	ClassSpellMask_HunterTraps   = ClassSpellMask_HunterExplosiveTrap | ClassSpellMask_HunterFreezingTrap | ClassSpellMask_HunterImmolationTrap
	ClassSpellMask_HunterShots   = ClassSpellMask_HunterAimedShot | ClassSpellMask_HunterArcaneShot | ClassSpellMask_HunterChimeraShot | ClassSpellMask_HunterExplosiveShot | ClassSpellMask_HunterKillShot | ClassSpellMask_HunterMultiShot | ClassSpellMask_HunterSteadyShot
	ClassSpellMask_HunterStrikes = ClassSpellMask_HunterFlankingStrike | ClassSpellMask_HunterRaptorStrike | ClassSpellMask_HunterRaptorStrikeHit | ClassSpellMask_HunterWyvernStrike | ClassSpellMask_HunterCarve | ClassSpellMask_HunterCarveHit
	ClassSpellMask_HunterStings  = ClassSpellMask_HunterSerpentSting | ClassSpellMask_HunterSoFSerpentSting
)

type Hunter struct {
	core.Character

	ClassSpellScaling float64

	Talents *proto.HunterTalents
	Options *proto.HunterOptions

	latency     time.Duration
	timeToWeave time.Duration

	// Pet          *HunterPet
	// StampedePet  []*HunterPet
	// DireBeastPet *HunterPet
	// Thunderhawks []*ThunderhawkPet

	// The most recent time at which moving could have started, for trap weaving.
	mayMoveAt time.Duration

	AspectOfTheHawk *core.Spell

	// Hunter spells
	SerpentSting         *core.Spell
	ArcaneShot           *core.Spell
	ExplosiveTrap        *core.Spell
	ExplosiveShot        *core.Spell
	ImprovedSerpentSting *core.Spell
	RapidFire            *core.Spell

	Shots []*core.Spell

	BestialWrathAura *core.Aura

	// Fake spells to encapsulate weaving logic.
	HuntersMarkSpell *core.Spell
}

func (hunter *Hunter) GetCharacter() *core.Character {
	return &hunter.Character
}

func (hunter *Hunter) GetHunter() *Hunter {
	return hunter
}

func RegisterHunter() {
	core.RegisterAgentFactory(
		proto.Player_Hunter{},
		proto.Spec_SpecHunter,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewHunter(character, options, options.GetHunter().Options.ClassOptions)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Hunter)
			if !ok {
				panic("Invalid spec value for Hunter!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewHunter(character *core.Character, options *proto.Player, hunterOptions *proto.HunterOptions) *Hunter {
	hunter := &Hunter{
		Character: *character,
		Talents:   &proto.HunterTalents{},
		Options:   hunterOptions,
	}

	core.FillTalentsProto(hunter.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)
	// focusPerSecond := 4.0

	// kindredSpritsBonusFocus := core.TernaryFloat64(hunter.Spec == proto.Spec_SpecBeastMasteryHunter, 20, 0)
	//hunter.EnableFocusBar(100+kindredSpritsBonusFocus, focusPerSecond, true, nil, true)

	hunter.PseudoStats.CanParry = true

	hunter.EnableManaBar()

	// Passive bonus (used to be from quiver).
	//hunter.PseudoStats.RangedSpeedMultiplier *= 1.15
	rangedWeapon := hunter.WeaponFromRanged(0)

	hunter.EnableAutoAttacks(hunter, core.AutoAttackOptions{
		Ranged: rangedWeapon,
		//ReplaceMHSwing:  hunter.TryRaptorStrike, //Todo: Might be weaving
		AutoSwingRanged: true,
		AutoSwingMelee:  false,
	})

	hunter.AutoAttacks.RangedConfig().ApplyEffects = func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
		baseDamage := hunter.RangedWeaponDamage(sim, spell.RangedAttackPower())

		result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)

		spell.WaitTravelTime(sim, func(sim *core.Simulation) {
			spell.DealDamage(sim, result)
		})
	}

	hunter.AddStatDependencies()

	// hunter.Pet = hunter.NewHunterPet()
	// hunter.StampedePet = make([]*HunterPet, 4)
	// for index := range 4 {
	// 	hunter.StampedePet[index] = hunter.NewStampedePet(index)
	// }

	// hunter.DireBeastPet = hunter.NewDireBeastPet()

	// // Add 10 just to be protected against weird good luck :)
	// hunter.Thunderhawks = make([]*ThunderhawkPet, 10)
	// for index := range 10 {
	// 	hunter.Thunderhawks[index] = hunter.NewThunderhawkPet(index)
	// }

	return hunter
}

func (hunter *Hunter) Initialize() {
	hunter.AutoAttacks.RangedConfig().CritMultiplier = hunter.DefaultMeleeCritMultiplier()

	hunter.registerArcaneShotSpell()

}

func (hunter *Hunter) GetBaseDamageFromCoeff(coeff float64) float64 {
	return coeff * hunter.ClassSpellScaling
}

func (hunter *Hunter) ApplyTalents() {
	// hunter.applyThrillOfTheHunt()
	// hunter.ApplyHotfixes()
	// hunter.addBloodthirstyGloves()

	// if hunter.Pet != nil {
	// 	hunter.Pet.ApplyTalents()
	// }

	// hunter.ApplyArmorSpecializationEffect(stats.Agility, proto.ArmorType_ArmorTypeMail, 86538)
}

func (hunter *Hunter) AddStatDependencies() {
	hunter.AddStatDependency(stats.Agility, stats.AttackPower, 2)
	hunter.AddStatDependency(stats.Agility, stats.RangedAttackPower, 2)
	hunter.AddStatDependency(stats.Agility, stats.PhysicalCritPercent, core.CritPerAgiMaxLevel[hunter.Class])
}

func (hunter *Hunter) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	// raidBuffs.TrueshotAura = true

	// // if hunter.Talents.FerociousInspiration && hunter.Options.PetType != proto.HunterOptions_PetNone {
	// // 	raidBuffs.FerociousInspiration = true
	// // }

	// if hunter.Options.PetType == proto.HunterOptions_CoreHound {
	// 	raidBuffs.Bloodlust = true
	// }
	// switch hunter.Options.PetType {
	// case proto.HunterOptions_CoreHound:
	// 	raidBuffs.Bloodlust = true

	// case proto.HunterOptions_ShaleSpider:
	// 	raidBuffs.EmbraceOfTheShaleSpider = true

	// case proto.HunterOptions_Wolf:
	// 	raidBuffs.FuriousHowl = true
	// case proto.HunterOptions_Devilsaur:
	// 	raidBuffs.TerrifyingRoar = true
	// case proto.HunterOptions_WaterStrider:
	// 	raidBuffs.StillWater = true
	// case proto.HunterOptions_Hyena:
	// 	raidBuffs.CacklingHowl = true
	// case proto.HunterOptions_Serpent:
	// 	raidBuffs.SerpentsSwiftness = true
	// case proto.HunterOptions_SporeBat:
	// 	raidBuffs.MindQuickening = true
	// case proto.HunterOptions_Cat:
	// 	raidBuffs.RoarOfCourage = true
	// case proto.HunterOptions_SpiritBeast:
	// 	raidBuffs.SpiritBeastBlessing = true
	// }
	// if hunter.Options.PetType == proto.HunterOptions_ShaleSpider {
	// 	raidBuffs.BlessingOfKings = true
	// }

	// if hunter.Options.PetType == proto.HunterOptions_Wolf || hunter.Options.PetType == proto.HunterOptions_Devilsaur {
	// 	raidBuffs.FuriousHowl = true
	// }

	// TODO: Fix this to work with the new talent system.
	//
	//	if hunter.Talents.HuntingParty {
	//		raidBuffs.HuntingParty = true
	//	}
}

func (hunter *Hunter) AddPartyBuffs(_ *proto.PartyBuffs) {
}

func (hunter *Hunter) Reset(_ *core.Simulation) {
	hunter.mayMoveAt = 0
}

func (hunter *Hunter) OnEncounterStart(sim *core.Simulation) {
}

const (
	HunterSpellFlagsNone int64 = 0
	SpellMaskSpellRanged int64 = 1 << iota
	HunterSpellAutoShot
	HunterSpellSteadyShot
	HunterSpellCobraShot
	HunterSpellArcaneShot
	HunterSpellKillCommand
	HunterSpellChimeraShot
	HunterSpellExplosiveShot
	HunterSpellExplosiveTrap
	HunterSpellBlackArrow
	HunterSpellMultiShot
	HunterSpellAimedShot
	HunterSpellSerpentSting
	HunterSpellKillShot
	HunterSpellRapidFire
	HunterSpellBestialWrath
	HunterPetFocusDump
	HunterPetDamage
	HunterPetBeastCleaveHit
	HunterSpellFervor
	HunterSpellDireBeast
	HunterSpellAMurderOfCrows
	HunterSpellLynxRush
	HunterSpellGlaiveToss
	HunterSpellBarrage
	HunterSpellPowershot
	HunterSpellsAll = HunterSpellSteadyShot | HunterSpellCobraShot |
		HunterSpellArcaneShot | HunterSpellKillCommand | HunterSpellChimeraShot | HunterSpellExplosiveShot |
		HunterSpellExplosiveTrap | HunterSpellBlackArrow | HunterSpellMultiShot | HunterSpellAimedShot |
		HunterSpellSerpentSting | HunterSpellKillShot | HunterSpellRapidFire | HunterSpellBestialWrath
	HunterSpellsTalents = HunterSpellFervor | HunterSpellDireBeast | HunterSpellAMurderOfCrows | HunterSpellLynxRush | HunterSpellGlaiveToss | HunterSpellPowershot | HunterSpellBarrage
)

// Agent is a generic way to access underlying hunter on any of the agents.
type HunterAgent interface {
	GetHunter() *Hunter
}
