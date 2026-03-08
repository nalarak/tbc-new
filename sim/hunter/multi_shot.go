package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (hunter *Hunter) registerMultiShotSpell() {
	hunter.MultiShot = hunter.RegisterSpell(core.SpellConfig{
		ActionID:       core.ActionID{SpellID: 27021},
		SpellSchool:    core.SpellSchoolPhysical,
		ProcMask:       core.ProcMaskRangedSpecial,
		ClassSpellMask: HunterSpellMultiShot,
		Flags:          core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		MissileSpeed: 30,
		MinRange:     core.MaxMeleeRange,
		MaxRange:     HunterBaseMaxRange,

		ManaCost: core.ManaCostOptions{
			FlatCost: 275,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: 1,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 10,
			},

			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				cast.CastTime = spell.CastTime()
				hunter.AutoAttacks.StopRangedUntil(sim, sim.CurrentTime+cast.CastTime)
			},

			CastTime: func(spell *core.Spell) time.Duration {
				return hunter.AutoAttacks.RangedSwingWindup()
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   hunter.DefaultMeleeCritMultiplier(),
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			baseDamage := spell.RangedAttackPower()*0.2 +
				hunter.AutoAttacks.Ranged().BaseDamage(sim) +
				205

			if hunter.TalonOfAlarAura.IsActive() {
				baseDamage += 40
			}

			spell.CalcAoeDamage(sim, baseDamage, spell.OutcomeRangedHitAndCrit)

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealBatchedAoeDamage(sim)
			})
		},
	})
}
