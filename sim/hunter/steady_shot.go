package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (hunter *Hunter) registerSteadyShotSpell() {
	hunter.SteadyShot = hunter.RegisterSpell(core.SpellConfig{
		ActionID:       core.ActionID{SpellID: 34120},
		SpellSchool:    core.SpellSchoolPhysical,
		ClassSpellMask: HunterSpellSteadyShot,
		ProcMask:       core.ProcMaskRangedSpecial,
		Flags:          core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		MissileSpeed: 40,
		MinRange:     core.MaxMeleeRange,
		MaxRange:     HunterBaseMaxRange,

		ManaCost: core.ManaCostOptions{
			FlatCost: 110,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
			IgnoreHaste: true,

			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				cast.CastTime = spell.CastTime()
				hunter.AutoAttacks.StopRangedUntil(sim, sim.CurrentTime+cast.CastTime)
			},

			CastTime: func(spell *core.Spell) time.Duration {
				return time.Duration(float64(spell.DefaultCast.CastTime) / hunter.TotalRangedHasteMultiplier())
			},
		},

		DamageMultiplier: 1,
		CritMultiplier:   hunter.DefaultMeleeCritMultiplier(),
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			weaponDamage := hunter.AutoAttacks.Ranged().BaseDamage(sim) - hunter.AmmoDamageBonus

			if hunter.Ranged().Enchant.EffectID == 2723 {
				weaponDamage -= 12
			}

			baseDamage := 0.2*spell.RangedAttackPower() +
				weaponDamage*2.8/hunter.AutoAttacks.Ranged().SwingSpeed +
				150

			if hunter.TalonOfAlarAura.IsActive() {
				baseDamage += 40
			}

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	})
}
