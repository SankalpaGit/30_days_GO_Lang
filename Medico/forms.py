from django import forms
from.models import RequestBlood, DonateBlood


class RequestBloodForm(forms.ModelForm):
    class Meta:
        model = RequestBlood
        fields = ['name', 'address', 'gender', 'dob', 'age', 'blood_group', 'phone', 'email']
    
    def save(self, commit=True):
        instance = super(RequestBloodForm, self).save(commit=False)
        if commit:
            instance.save()
        return instance
    
class DonateBloodForm(forms.ModelForm):
    class Meta:
        model = DonateBlood
        fields = ['name', 'address', 'gender', 'dob', 'age', 'blood_group', 'phone', 'email']
    
    def save(self, commit=True):
        instance = super(DonateBloodForm, self).save(commit=False)
        if commit:
            instance.save()
        return instance


