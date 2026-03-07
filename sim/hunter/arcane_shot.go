package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (hunter *Hunter) registerArcaneShotSpell() {
	baseCost := 230.0

	hunter.ArcaneShot = hunter.RegisterSpell(core.SpellConfig{
		ClassSpellMask: HunterSpellArcaneShot,
		ActionID:       core.ActionID{SpellID: 27019},
		SpellSchool:    core.SpellSchoolArcane,

		ProcMask: core.ProcMaskRangedSpecial,
		Flags:    core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		MissileSpeed: 24,

		ManaCost: core.ManaCostOptions{
			FlatCost:        int32(baseCost),
			PercentModifier: 1,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.02*float64(hunter.Talents.Efficiency)),
				GCD:  core.GCDDefault + hunter.latency,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second*6 - time.Millisecond*200*time.Duration(hunter.Talents.ImprovedArcaneShot),
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			wepDmg := hunter.AutoAttacks.Ranged().CalculateNormalizedWeaponDamage(sim, spell.RangedAttackPower())

			baseDamage := wepDmg + (hunter.ClassSpellScaling*0.15 + 273)

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeRangedHitAndCrit)

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	})

}
