# public_portofolio
a showcase of past and present personal projects or projects from previous work ---- a portofolio of sorts

class Attributes(DonKane):
	@staticmethod
	def contact() -> Donkane:
	    proton   = "don24kane@proton.me"
	    return proton
	def life() -> tuple:
		langs         = ['Romanian', 'English']
		nationalities = self.langs.remove('Romanian')
		age           = 23
		return langs, nationalities, age
	def coding() -> tuple:
		langs = {
			'advanced':   ['c++', 'js/ts', 'python'],
			'intermediate': ['go', 'c#'],
      'begginer': ['java', 'c'],
			'learning': ['carbon', 'php']
		}
		specialities  = ['soft engineer', 'web/app engineer & tester', 'fullstack, pref back-end']
		environnement = ['vscode', '.sh', 'notepad', 'notepad++']
		return langs, specialities, environnement
